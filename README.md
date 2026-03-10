**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-10 03:15:35 UTC（2026-03-10 11:15:35 UTC+8）

**代理总数：68**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 68 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 154.3.236.202:3128 | ✓ 999ms | 否 | ✓ 1230ms | ✓ 1593ms | ✓ 1142ms | http |
| 1.231.81.166:3128 | ✓ 1435ms | ✓ 1609ms | ✓ 980ms | ✓ 1087ms | ✓ 1097ms | http |
| 91.107.141.42:8081 | ✓ 688ms | 否 | ✓ 1582ms | 否 | ✓ 1646ms | http |
| 168.235.110.63:3128 | ✓ 1325ms | ✓ 1632ms | ✓ 1802ms | ✓ 1325ms | ✓ 967ms | http |
| 89.185.85.138:1080 | ✓ 1001ms | ✓ 1854ms | ✓ 1280ms | ✓ 1898ms | ✓ 1318ms | http |
| 165.227.5.10:8888 | 否 | 否 | ✓ 1234ms | ✓ 851ms | ✓ 1593ms | http |
| 45.136.131.47:8443 | ✓ 101ms | ✓ 1181ms | ✓ 88ms | ✓ 695ms | ✓ 1117ms | http |
| 20.210.76.175:8561 | ✓ 1103ms | ✓ 912ms | ✓ 620ms | ✓ 804ms | ✓ 636ms | http |
| 20.78.26.206:8561 | ✓ 1077ms | ✓ 873ms | ✓ 686ms | ✓ 820ms | ✓ 598ms | http |
| 20.210.76.178:8561 | ✓ 1541ms | ✓ 901ms | ✓ 453ms | ✓ 771ms | ✓ 599ms | http |
| 35.225.22.61:80 | ✓ 446ms | 否 | ✓ 943ms | ✓ 1146ms | ✓ 827ms | http |
| 20.27.15.49:8561 | ✓ 1077ms | ✓ 1892ms | ✓ 483ms | ✓ 755ms | ✓ 623ms | http |
| 20.78.118.91:8561 | ✓ 1597ms | 否 | ✓ 452ms | ✓ 779ms | ✓ 697ms | http |
| 20.210.76.104:8561 | ✓ 1541ms | ✓ 1279ms | ✓ 445ms | ✓ 758ms | ✓ 640ms | http |
| 95.3.9.78:3128 | ✓ 862ms | ✓ 1910ms | ✓ 848ms | ✓ 1796ms | ✓ 1387ms | http |
| 61.72.110.114:3128 | ✓ 1152ms | ✓ 1331ms | 否 | ✓ 1101ms | ✓ 1960ms | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1836ms | ✓ 1149ms | ✓ 1034ms | http |
| 61.72.110.54:3128 | 否 | 否 | ✓ 1111ms | ✓ 1006ms | ✓ 1617ms | http |
| 202.155.12.161:443 | ✓ 1652ms | 否 | ✓ 1528ms | 否 | ✓ 1071ms | http |
| 46.183.25.8:443 | ✓ 734ms | 否 | ✓ 1714ms | ✓ 1776ms | 否 | http |
| 39.104.201.40:7890 | ✓ 926ms | ✓ 1189ms | ✓ 959ms | ✓ 1220ms | ✓ 968ms | http |
| 101.43.255.96:80 | ✓ 1014ms | 否 | ✓ 1011ms | ✓ 1644ms | ✓ 1306ms | http |
| 81.70.169.194:80 | ✓ 1256ms | ✓ 1321ms | 否 | ✓ 1579ms | ✓ 1004ms | http |
| 190.9.109.198:999 | ✓ 1136ms | ✓ 1643ms | ✓ 1097ms | ✓ 1440ms | ✓ 1386ms | http |
| 190.9.109.207:999 | ✓ 1008ms | 否 | ✓ 1203ms | ✓ 1342ms | ✓ 1162ms | http |
| 101.47.73.135:3128 | ✓ 1284ms | 否 | ✓ 825ms | ✓ 1471ms | 否 | http |
| 95.3.9.78:8080 | ✓ 1096ms | 否 | ✓ 1226ms | 否 | ✓ 1420ms | http |
| 185.191.236.162:3128 | ✓ 1550ms | 否 | ✓ 1603ms | 否 | ✓ 1203ms | http |
| 115.231.181.40:8128 | ✓ 1533ms | ✓ 1084ms | ✓ 857ms | ✓ 1167ms | ✓ 1088ms | http |
| 121.237.181.137:8888 | 否 | ✓ 1091ms | ✓ 906ms | ✓ 1150ms | ✓ 998ms | http |
| 193.168.173.136:443 | ✓ 1488ms | 否 | ✓ 905ms | ✓ 1911ms | 否 | http |
| 194.213.18.200:443 | ✓ 1782ms | ✓ 1621ms | ✓ 1291ms | ✓ 1706ms | 否 | http |
| 120.92.212.16:7890 | ✓ 986ms | ✓ 1260ms | 否 | 否 | ✓ 994ms | http |
| 138.124.53.25:7443 | ✓ 1433ms | 否 | ✓ 1961ms | ✓ 1938ms | ✓ 1779ms | http |
| 116.80.82.227:3172 | ✓ 1835ms | 否 | 否 | ✓ 1823ms | ✓ 1631ms | http |
| 106.14.203.63:3333 | ✓ 842ms | ✓ 1091ms | ✓ 1277ms | ✓ 1771ms | ✓ 918ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 905ms | ✓ 934ms | ✓ 811ms | http |
| 125.128.12.14:3128 | ✓ 1581ms | ✓ 1635ms | ✓ 1277ms | ✓ 963ms | ✓ 775ms | http |
| 14.225.222.213:7890 | ✓ 1463ms | 否 | 否 | ✓ 1341ms | ✓ 907ms | http |
| 116.80.96.101:3172 | 否 | 否 | ✓ 1509ms | ✓ 1826ms | ✓ 1893ms | http |
| 34.101.184.164:3128 | ✓ 1769ms | 否 | 否 | ✓ 1372ms | ✓ 1137ms | http |
| 207.5.207.47:3128 | 否 | 否 | ✓ 1436ms | ✓ 1062ms | ✓ 859ms | http |
| 103.149.99.128:3128 | 否 | 否 | ✓ 1547ms | ✓ 1230ms | ✓ 947ms | http |
| 120.92.212.16:8890 | ✓ 997ms | ✓ 1366ms | 否 | ✓ 1215ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1495ms | ✓ 1626ms | ✓ 1236ms | ✓ 1054ms | ✓ 1569ms | http |
| 45.136.130.223:8443 | ✓ 641ms | ✓ 600ms | ✓ 252ms | ✓ 679ms | ✓ 506ms | http |
| 125.128.12.144:3128 | 否 | 否 | ✓ 1510ms | ✓ 1090ms | ✓ 840ms | http |
| 14.225.212.37:7890 | ✓ 1968ms | 否 | 否 | ✓ 1860ms | ✓ 836ms | http |
| 45.77.246.231:80 | ✓ 1399ms | 否 | ✓ 1699ms | ✓ 1052ms | ✓ 1697ms | http |
| 103.134.246.42:3172 | ✓ 1458ms | 否 | ✓ 1398ms | ✓ 1552ms | 否 | http |
| 1.234.153.14:80 | ✓ 621ms | ✓ 904ms | ✓ 726ms | ✓ 1256ms | ✓ 1374ms | http |
| 45.140.147.155:1081 | ✓ 1845ms | ✓ 1552ms | ✓ 639ms | 否 | ✓ 1174ms | http |
| 45.88.0.115:3128 | ✓ 1073ms | ✓ 1965ms | ✓ 1701ms | 否 | ✓ 1344ms | http |
| 45.129.141.143:3128 | ✓ 1312ms | 否 | ✓ 1640ms | 否 | ✓ 1602ms | http |
| 103.39.51.190:8080 | ✓ 1863ms | 否 | ✓ 1850ms | ✓ 1590ms | ✓ 1468ms | http |
| 45.186.6.104:3128 | ✓ 1226ms | ✓ 1971ms | ✓ 1826ms | 否 | 否 | http |
| 178.236.245.59:3128 | ✓ 1003ms | 否 | ✓ 1208ms | 否 | ✓ 1734ms | http |
| 136.49.34.18:8888 | ✓ 1556ms | ✓ 1313ms | ✓ 709ms | ✓ 1118ms | ✓ 916ms | http |
| 178.236.245.17:3128 | ✓ 730ms | ✓ 1729ms | ✓ 800ms | ✓ 1881ms | ✓ 1437ms | http |
| 107.172.125.217:3128 | ✓ 930ms | 否 | ✓ 760ms | ✓ 694ms | ✓ 523ms | http |
| 103.139.138.194:3128 | ✓ 1081ms | 否 | ✓ 1755ms | ✓ 1421ms | ✓ 1315ms | http |
| 152.70.98.46:8888 | ✓ 1439ms | 否 | ✓ 1354ms | ✓ 1345ms | ✓ 1915ms | http |
| 61.52.131.172:8443 | ✓ 921ms | 否 | ✓ 938ms | ✓ 1152ms | ✓ 1583ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1389ms | ✓ 1350ms | ✓ 1266ms | http |
| 45.136.198.40:3128 | ✓ 1331ms | ✓ 1767ms | 否 | 否 | ✓ 1731ms | http |
| 47.77.193.180:1080 | ✓ 676ms | ✓ 1453ms | ✓ 403ms | ✓ 732ms | ✓ 554ms | http |
| 128.199.113.85:9090 | ✓ 924ms | 否 | ✓ 1406ms | ✓ 1546ms | ✓ 1380ms | http |
| 222.107.27.7:8076 | ✓ 1178ms | 否 | ✓ 1882ms | ✓ 1309ms | ✓ 1190ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
