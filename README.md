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

最后更新：2026-03-07 22:17:37 UTC（2026-03-08 06:17:37 UTC+8）

**代理总数：70**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 70 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 347ms | 否 | 否 | ✓ 1329ms | ✓ 1160ms | http |
| 152.42.213.210:8080 | ✓ 751ms | 否 | 否 | ✓ 1093ms | ✓ 839ms | http |
| 35.225.22.61:80 | ✓ 900ms | ✓ 1398ms | 否 | ✓ 944ms | 否 | http |
| 1.231.81.166:3128 | ✓ 820ms | ✓ 1047ms | ✓ 1031ms | ✓ 984ms | ✓ 786ms | http |
| 103.84.95.54:7890 | ✓ 785ms | 否 | 否 | ✓ 1049ms | ✓ 724ms | http |
| 178.236.245.17:3128 | ✓ 1015ms | 否 | ✓ 1139ms | 否 | ✓ 1662ms | http |
| 89.185.85.138:1080 | ✓ 1014ms | 否 | ✓ 1273ms | ✓ 1862ms | 否 | http |
| 178.236.245.59:3128 | ✓ 982ms | 否 | ✓ 1141ms | 否 | ✓ 1682ms | http |
| 121.128.121.54:3128 | ✓ 1300ms | ✓ 1440ms | 否 | 否 | ✓ 802ms | http |
| 217.76.245.80:999 | ✓ 879ms | 否 | ✓ 1126ms | ✓ 1490ms | ✓ 1281ms | http |
| 61.72.221.194:3128 | ✓ 1221ms | ✓ 1442ms | 否 | ✓ 1078ms | ✓ 989ms | http |
| 115.231.181.40:8128 | ✓ 889ms | ✓ 1071ms | ✓ 888ms | ✓ 1163ms | ✓ 942ms | http |
| 4.213.222.169:3128 | ✓ 1249ms | 否 | ✓ 1323ms | ✓ 1611ms | 否 | http |
| 138.124.53.25:7443 | ✓ 808ms | ✓ 1945ms | 否 | ✓ 1711ms | ✓ 1290ms | http |
| 14.225.217.30:7890 | ✓ 1342ms | ✓ 1397ms | ✓ 1971ms | 否 | ✓ 1961ms | http |
| 46.183.25.8:443 | ✓ 1445ms | 否 | 否 | ✓ 1980ms | ✓ 1501ms | http |
| 61.72.221.94:3128 | ✓ 656ms | 否 | ✓ 1571ms | 否 | ✓ 855ms | http |
| 202.155.12.161:443 | ✓ 1691ms | ✓ 1508ms | 否 | ✓ 1201ms | ✓ 1186ms | http |
| 101.43.255.96:80 | ✓ 1012ms | ✓ 1272ms | ✓ 983ms | ✓ 1349ms | ✓ 970ms | http |
| 81.70.169.194:80 | ✓ 1029ms | ✓ 1286ms | ✓ 1073ms | ✓ 1314ms | ✓ 1002ms | http |
| 45.186.6.104:3128 | ✓ 1107ms | ✓ 1890ms | ✓ 1672ms | 否 | 否 | http |
| 222.228.171.92:8080 | ✓ 1752ms | 否 | 否 | ✓ 1842ms | ✓ 1454ms | http |
| 91.193.240.157:9877 | ✓ 1486ms | 否 | ✓ 1762ms | 否 | ✓ 1717ms | http |
| 47.101.159.19:8899 | ✓ 881ms | ✓ 1063ms | ✓ 889ms | ✓ 1108ms | ✓ 887ms | http |
| 46.249.103.192:443 | ✓ 807ms | 否 | ✓ 1168ms | ✓ 1888ms | 否 | http |
| 39.104.201.40:7890 | ✓ 947ms | ✓ 1221ms | ✓ 1014ms | ✓ 1288ms | ✓ 951ms | http |
| 103.215.36.88:19698 | ✓ 1060ms | ✓ 1358ms | ✓ 1229ms | ✓ 1409ms | ✓ 1063ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1065ms | 否 | ✓ 1941ms | ✓ 950ms | http |
| 190.9.109.198:999 | ✓ 781ms | ✓ 1523ms | ✓ 1283ms | ✓ 1754ms | ✓ 1200ms | http |
| 190.9.109.199:999 | ✓ 837ms | ✓ 1594ms | ✓ 1238ms | ✓ 1584ms | 否 | http |
| 67.169.98.211:443 | ✓ 1380ms | ✓ 1784ms | 否 | ✓ 1222ms | 否 | http |
| 194.59.204.87:9080 | ✓ 639ms | ✓ 1582ms | ✓ 615ms | 否 | 否 | http |
| 61.72.221.234:3128 | 否 | 否 | ✓ 870ms | ✓ 1077ms | ✓ 847ms | http |
| 85.9.195.140:1080 | ✓ 1186ms | 否 | ✓ 1600ms | ✓ 1746ms | ✓ 1248ms | http |
| 193.228.139.78:8888 | ✓ 636ms | ✓ 1775ms | ✓ 1926ms | 否 | ✓ 1371ms | http |
| 124.121.2.131:8080 | 否 | 否 | ✓ 1695ms | ✓ 1577ms | ✓ 1560ms | http |
| 120.92.212.16:7890 | ✓ 1337ms | 否 | ✓ 942ms | ✓ 1484ms | ✓ 1747ms | http |
| 125.128.12.144:3128 | ✓ 1549ms | ✓ 1262ms | ✓ 981ms | ✓ 1073ms | ✓ 1022ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1508ms | 否 | ✓ 1263ms | ✓ 1499ms | http |
| 159.89.31.62:8080 | ✓ 1637ms | ✓ 1805ms | 否 | 否 | ✓ 1674ms | http |
| 125.128.12.14:3128 | ✓ 1529ms | ✓ 1854ms | ✓ 1587ms | ✓ 1527ms | ✓ 1755ms | http |
| 61.72.110.54:3128 | 否 | ✓ 1962ms | 否 | ✓ 1105ms | ✓ 854ms | http |
| 14.56.107.244:3128 | ✓ 1444ms | ✓ 1721ms | ✓ 1988ms | ✓ 1639ms | ✓ 1875ms | http |
| 117.72.46.30:7890 | 否 | ✓ 1222ms | ✓ 1207ms | ✓ 1273ms | ✓ 1069ms | http |
| 185.243.218.43:49153 | 否 | ✓ 1971ms | ✓ 1590ms | ✓ 1930ms | ✓ 1442ms | http |
| 101.32.244.83:8080 | ✓ 1336ms | 否 | ✓ 1040ms | ✓ 1259ms | ✓ 1231ms | http |
| 121.43.196.210:8222 | ✓ 949ms | ✓ 1095ms | ✓ 870ms | ✓ 1151ms | ✓ 875ms | http |
| 121.43.196.213:8222 | ✓ 990ms | ✓ 1106ms | ✓ 845ms | ✓ 1169ms | ✓ 918ms | http |
| 114.55.226.123:10086 | ✓ 1129ms | ✓ 1635ms | ✓ 991ms | ✓ 1346ms | ✓ 1069ms | http |
| 168.235.110.63:3128 | ✓ 1437ms | ✓ 1335ms | ✓ 1179ms | 否 | ✓ 1103ms | http |
| 165.227.5.10:8888 | ✓ 626ms | ✓ 999ms | ✓ 769ms | 否 | ✓ 705ms | http |
| 162.240.154.26:3128 | 否 | 否 | ✓ 914ms | ✓ 992ms | ✓ 1555ms | http |
| 210.223.44.230:3128 | ✓ 1850ms | 否 | ✓ 1065ms | ✓ 915ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1786ms | ✓ 1452ms | ✓ 1701ms | ✓ 1811ms | ✓ 1102ms | http |
| 8.217.129.162:1080 | 否 | 否 | ✓ 1707ms | ✓ 888ms | ✓ 1181ms | http |
| 129.213.162.27:17777 | ✓ 886ms | ✓ 1187ms | 否 | 否 | ✓ 1329ms | http |
| 192.73.243.98:3128 | ✓ 707ms | 否 | ✓ 318ms | ✓ 1090ms | ✓ 831ms | http |
| 51.158.61.240:13128 | ✓ 1127ms | 否 | ✓ 1234ms | ✓ 1467ms | ✓ 1292ms | http |
| 94.156.112.198:3128 | ✓ 571ms | ✓ 1592ms | ✓ 1681ms | ✓ 1971ms | ✓ 1659ms | http |
| 103.151.227.134:3125 | ✓ 1375ms | 否 | ✓ 1512ms | ✓ 1426ms | ✓ 1402ms | http |
| 161.248.158.142:3125 | ✓ 1374ms | 否 | ✓ 1344ms | ✓ 1685ms | ✓ 1487ms | http |
| 59.8.203.55:80 | ✓ 1714ms | ✓ 1574ms | ✓ 873ms | ✓ 976ms | ✓ 792ms | http |
| 38.180.2.107:3128 | ✓ 901ms | ✓ 1823ms | ✓ 1739ms | 否 | ✓ 1942ms | http |
| 62.113.119.14:8080 | ✓ 1964ms | ✓ 1587ms | ✓ 912ms | ✓ 1776ms | ✓ 1179ms | http |
| 172.212.68.37:3128 | 否 | 否 | ✓ 889ms | ✓ 1734ms | ✓ 1180ms | http |
| 200.125.171.254:999 | ✓ 754ms | 否 | ✓ 1214ms | ✓ 1447ms | ✓ 1545ms | http |
| 1.225.116.115:1080 | ✓ 1601ms | ✓ 1320ms | ✓ 1301ms | ✓ 1024ms | ✓ 923ms | http |
| 45.136.198.40:3128 | ✓ 802ms | ✓ 1922ms | 否 | 否 | ✓ 1912ms | http |
| 103.39.51.190:8080 | ✓ 1840ms | 否 | 否 | ✓ 1386ms | ✓ 1593ms | http |
| 103.145.34.67:8080 | 否 | 否 | ✓ 1279ms | ✓ 1454ms | ✓ 1463ms | http |

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
