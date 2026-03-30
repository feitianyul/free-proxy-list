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

最后更新：2026-03-30 17:53:53 UTC（2026-03-31 01:53:53 UTC+8）

**代理总数：69**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 69 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 39.185.46.193:5911 | ✓ 803ms | ✓ 842ms | ✓ 698ms | ✓ 1012ms | ✓ 674ms | http |
| 147.161.210.140:8800 | ✓ 730ms | 否 | ✓ 843ms | ✓ 965ms | ✓ 869ms | http |
| 43.99.54.236:5555 | ✓ 665ms | 否 | 否 | ✓ 992ms | ✓ 1403ms | http |
| 167.103.115.102:8800 | ✓ 1585ms | 否 | ✓ 1124ms | ✓ 1029ms | ✓ 1220ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1342ms | ✓ 1043ms | ✓ 1281ms | ✓ 1053ms | http |
| 167.103.34.108:8800 | ✓ 1636ms | 否 | ✓ 1164ms | ✓ 1597ms | ✓ 1674ms | http |
| 183.249.5.105:22222 | ✓ 663ms | ✓ 864ms | ✓ 647ms | ✓ 879ms | ✓ 684ms | http |
| 183.249.5.117:22222 | ✓ 641ms | ✓ 1062ms | ✓ 658ms | ✓ 1021ms | ✓ 702ms | http |
| 35.225.22.61:80 | ✓ 611ms | 否 | ✓ 549ms | ✓ 1476ms | ✓ 970ms | http |
| 1.231.81.166:3128 | ✓ 1612ms | 否 | ✓ 1319ms | ✓ 864ms | ✓ 952ms | http |
| 120.92.212.16:7890 | ✓ 880ms | ✓ 1524ms | ✓ 1102ms | 否 | ✓ 935ms | http |
| 222.184.48.251:22222 | ✓ 1608ms | 否 | ✓ 1614ms | 否 | ✓ 1527ms | http |
| 167.103.31.122:8800 | ✓ 1902ms | 否 | ✓ 1904ms | 否 | ✓ 1849ms | http |
| 207.254.71.62:8088 | ✓ 820ms | ✓ 1927ms | ✓ 1667ms | 否 | ✓ 1962ms | http |
| 167.103.144.127:8800 | ✓ 1308ms | 否 | ✓ 968ms | ✓ 1203ms | ✓ 1092ms | http |
| 95.213.217.168:52004 | ✓ 1347ms | 否 | ✓ 1012ms | ✓ 1895ms | ✓ 1424ms | http |
| 147.161.239.240:8800 | ✓ 1303ms | 否 | ✓ 1352ms | ✓ 1910ms | ✓ 1529ms | http |
| 101.47.73.135:3128 | ✓ 1305ms | 否 | ✓ 1439ms | ✓ 1409ms | ✓ 1649ms | http |
| 222.184.48.242:22222 | 否 | ✓ 1244ms | ✓ 933ms | ✓ 1195ms | ✓ 1794ms | http |
| 180.250.219.58:53281 | ✓ 1848ms | 否 | ✓ 1789ms | ✓ 1932ms | ✓ 1960ms | http |
| 103.84.95.54:7890 | ✓ 867ms | 否 | ✓ 774ms | 否 | ✓ 1244ms | http |
| 101.43.127.100:8877 | ✓ 1596ms | ✓ 1062ms | ✓ 848ms | ✓ 1201ms | ✓ 857ms | http |
| 120.92.212.16:8890 | ✓ 920ms | ✓ 1385ms | 否 | ✓ 1693ms | ✓ 1142ms | http |
| 31.192.106.135:8010 | ✓ 1547ms | 否 | 否 | ✓ 1910ms | ✓ 1761ms | http |
| 5.104.87.17:8051 | ✓ 1453ms | 否 | ✓ 1911ms | ✓ 1815ms | ✓ 1425ms | http |
| 116.80.49.165:3172 | ✓ 1602ms | 否 | 否 | ✓ 1964ms | ✓ 1700ms | http |
| 38.34.179.27:8452 | ✓ 618ms | ✓ 1255ms | 否 | ✓ 820ms | ✓ 1329ms | http |
| 42.96.16.158:1311 | ✓ 1372ms | 否 | ✓ 1222ms | ✓ 1760ms | ✓ 1150ms | http |
| 16.78.119.130:443 | ✓ 1848ms | 否 | ✓ 1974ms | 否 | ✓ 1919ms | http |
| 183.249.5.111:22222 | ✓ 785ms | ✓ 965ms | ✓ 660ms | ✓ 1266ms | ✓ 695ms | http |
| 183.249.5.110:22222 | ✓ 744ms | ✓ 889ms | ✓ 941ms | ✓ 995ms | ✓ 1871ms | http |
| 222.184.48.252:22222 | ✓ 937ms | ✓ 1142ms | ✓ 1356ms | ✓ 1308ms | ✓ 950ms | http |
| 209.126.84.232:8888 | ✓ 1106ms | 否 | ✓ 1264ms | ✓ 1274ms | ✓ 1413ms | http |
| 103.82.23.118:5234 | ✓ 1422ms | 否 | ✓ 1237ms | ✓ 1601ms | ✓ 1710ms | http |
| 177.234.217.88:999 | ✓ 1073ms | ✓ 1842ms | ✓ 1094ms | ✓ 1858ms | ✓ 1583ms | http |
| 121.230.9.185:1080 | ✓ 1451ms | 否 | ✓ 1366ms | ✓ 1697ms | ✓ 1020ms | http |
| 8.219.97.248:80 | ✓ 1812ms | 否 | 否 | ✓ 1672ms | ✓ 1213ms | http |
| 31.56.177.74:1090 | ✓ 994ms | ✓ 1994ms | ✓ 1647ms | 否 | ✓ 1704ms | http |
| 45.12.151.226:2829 | ✓ 1224ms | 否 | 否 | ✓ 1903ms | ✓ 1477ms | http |
| 115.231.181.40:8128 | ✓ 1098ms | 否 | ✓ 1986ms | ✓ 1140ms | 否 | http |
| 45.136.198.40:3128 | ✓ 1759ms | ✓ 1779ms | 否 | 否 | ✓ 1880ms | http |
| 59.11.138.229:3128 | ✓ 1328ms | ✓ 953ms | ✓ 976ms | ✓ 970ms | ✓ 753ms | http |
| 106.117.208.101:7890 | ✓ 990ms | ✓ 1263ms | ✓ 985ms | ✓ 1267ms | ✓ 1274ms | http |
| 125.64.244.100:8880 | ✓ 1582ms | ✓ 1623ms | ✓ 1447ms | ✓ 1814ms | ✓ 1527ms | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 1432ms | ✓ 1540ms | ✓ 1221ms | http |
| 45.140.147.82:1081 | ✓ 623ms | ✓ 1547ms | ✓ 1456ms | ✓ 1757ms | ✓ 1592ms | http |
| 121.230.9.205:1080 | 否 | ✓ 1682ms | ✓ 1126ms | ✓ 1535ms | ✓ 1101ms | http |
| 59.46.216.131:30001 | ✓ 928ms | 否 | ✓ 1021ms | 否 | ✓ 1018ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1185ms | ✓ 1409ms | ✓ 1116ms | http |
| 103.178.86.178:8080 | ✓ 1466ms | 否 | ✓ 1699ms | ✓ 1556ms | ✓ 1255ms | http |
| 47.105.98.23:3128 | 否 | ✓ 1613ms | ✓ 1123ms | 否 | ✓ 1650ms | http |
| 45.174.241.241:999 | 否 | 否 | ✓ 1760ms | ✓ 1649ms | ✓ 1466ms | http |
| 150.249.255.91:3128 | ✓ 1301ms | ✓ 882ms | ✓ 610ms | ✓ 786ms | ✓ 598ms | http |
| 103.39.51.190:8080 | ✓ 1845ms | 否 | 否 | ✓ 1747ms | ✓ 1556ms | http |
| 103.113.70.189:1081 | ✓ 331ms | 否 | ✓ 845ms | 否 | ✓ 1925ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1087ms | ✓ 1269ms | ✓ 1821ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1235ms | ✓ 1460ms | ✓ 1191ms | http |
| 167.71.196.28:8080 | ✓ 695ms | 否 | ✓ 1135ms | ✓ 1022ms | 否 | http |
| 62.113.119.14:8080 | ✓ 877ms | 否 | ✓ 1104ms | 否 | ✓ 1475ms | http |
| 107.173.52.221:3128 | ✓ 1106ms | 否 | ✓ 1391ms | ✓ 1425ms | ✓ 1212ms | http |
| 205.164.46.6:3129 | ✓ 848ms | ✓ 1297ms | ✓ 1205ms | 否 | 否 | http |
| 38.34.183.11:8451 | ✓ 531ms | ✓ 771ms | ✓ 290ms | ✓ 734ms | ✓ 1557ms | http |
| 210.223.44.230:3128 | ✓ 810ms | ✓ 1403ms | ✓ 866ms | 否 | ✓ 720ms | http |
| 116.80.65.81:3172 | ✓ 1452ms | ✓ 1960ms | ✓ 1599ms | 否 | ✓ 1631ms | http |
| 208.87.243.199:7878 | ✓ 1160ms | 否 | ✓ 1482ms | ✓ 1612ms | 否 | http |
| 34.96.238.40:8080 | ✓ 1062ms | 否 | 否 | ✓ 1478ms | ✓ 1543ms | http |
| 164.90.206.15:3128 | ✓ 769ms | ✓ 1900ms | ✓ 1430ms | ✓ 1776ms | ✓ 1600ms | http |
| 194.59.204.87:9080 | ✓ 1387ms | 否 | ✓ 1772ms | 否 | ✓ 1835ms | http |
| 222.228.171.92:8080 | ✓ 687ms | 否 | ✓ 1193ms | ✓ 1883ms | ✓ 1006ms | http |

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
