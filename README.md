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

最后更新：2026-04-06 11:58:04 UTC（2026-04-06 19:58:04 UTC+8）

**代理总数：83**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 681ms | 否 | ✓ 1003ms | ✓ 1173ms | ✓ 921ms | http |
| 4.216.195.194:3128 | ✓ 653ms | 否 | ✓ 1264ms | ✓ 1102ms | ✓ 936ms | http |
| 113.160.132.26:8080 | ✓ 1754ms | ✓ 1367ms | ✓ 1195ms | ✓ 1205ms | ✓ 980ms | http |
| 159.223.71.162:443 | ✓ 1699ms | 否 | ✓ 1308ms | ✓ 1115ms | ✓ 926ms | http |
| 159.223.71.162:8080 | ✓ 1693ms | 否 | ✓ 1314ms | ✓ 1109ms | ✓ 947ms | http |
| 167.103.115.102:8800 | ✓ 1612ms | ✓ 1913ms | ✓ 1300ms | 否 | ✓ 1286ms | http |
| 150.241.71.15:1080 | ✓ 996ms | 否 | 否 | ✓ 1729ms | ✓ 1639ms | http |
| 111.227.254.10:22222 | ✓ 1081ms | ✓ 1360ms | 否 | ✓ 1377ms | 否 | http |
| 111.227.254.9:22222 | 否 | ✓ 1242ms | ✓ 1111ms | ✓ 1958ms | 否 | http |
| 167.103.144.127:8800 | ✓ 995ms | 否 | ✓ 1350ms | ✓ 1456ms | ✓ 1381ms | http |
| 167.103.34.108:8800 | ✓ 1513ms | 否 | ✓ 1660ms | ✓ 1442ms | ✓ 1561ms | http |
| 111.227.254.12:22222 | 否 | ✓ 1662ms | ✓ 1437ms | 否 | ✓ 1447ms | http |
| 111.227.254.11:22222 | ✓ 1044ms | 否 | ✓ 980ms | ✓ 1403ms | ✓ 1096ms | http |
| 45.167.125.21:999 | ✓ 1253ms | 否 | 否 | ✓ 1959ms | ✓ 1540ms | http |
| 103.166.185.54:3128 | ✓ 1735ms | ✓ 1531ms | ✓ 1141ms | ✓ 1283ms | ✓ 1292ms | http |
| 167.103.31.122:8800 | ✓ 1427ms | 否 | ✓ 1388ms | ✓ 1690ms | ✓ 1544ms | http |
| 35.225.22.61:80 | ✓ 898ms | ✓ 1437ms | 否 | 否 | ✓ 1107ms | http |
| 45.149.92.147:5001 | ✓ 673ms | 否 | ✓ 812ms | ✓ 841ms | ✓ 837ms | http |
| 1.231.81.166:3128 | ✓ 1532ms | ✓ 1482ms | ✓ 1026ms | ✓ 1469ms | ✓ 1451ms | http |
| 147.161.239.240:8800 | ✓ 1180ms | 否 | ✓ 1586ms | ✓ 1724ms | ✓ 1560ms | http |
| 181.78.44.63:999 | ✓ 903ms | 否 | ✓ 1445ms | ✓ 1757ms | ✓ 1275ms | http |
| 180.250.219.58:53281 | 否 | 否 | ✓ 1687ms | ✓ 1899ms | ✓ 1828ms | http |
| 46.39.105.157:8080 | ✓ 1846ms | 否 | ✓ 1305ms | ✓ 1749ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1397ms | 否 | ✓ 1103ms | 否 | ✓ 1200ms | http |
| 59.46.216.131:30001 | ✓ 949ms | 否 | ✓ 1045ms | ✓ 1656ms | ✓ 1990ms | http |
| 20.118.221.52:3128 | ✓ 1119ms | ✓ 1355ms | ✓ 863ms | ✓ 1230ms | ✓ 916ms | http |
| 117.86.6.35:1080 | ✓ 967ms | ✓ 1735ms | ✓ 1961ms | ✓ 1592ms | ✓ 997ms | http |
| 34.101.184.164:3128 | ✓ 1667ms | 否 | ✓ 1704ms | ✓ 1280ms | ✓ 998ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1085ms | 否 | ✓ 1429ms | ✓ 1503ms | http |
| 101.43.127.100:8877 | ✓ 922ms | ✓ 1141ms | ✓ 871ms | ✓ 1091ms | 否 | http |
| 43.167.237.94:3128 | ✓ 1683ms | 否 | ✓ 1334ms | ✓ 1173ms | ✓ 1472ms | http |
| 209.38.154.7:1080 | ✓ 1330ms | 否 | ✓ 693ms | 否 | ✓ 544ms | http |
| 62.113.119.14:8080 | ✓ 1368ms | ✓ 1940ms | ✓ 1248ms | ✓ 1791ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1507ms | ✓ 1044ms | ✓ 1034ms | 否 | ✓ 1008ms | http |
| 218.108.131.186:17890 | ✓ 821ms | ✓ 1001ms | ✓ 865ms | ✓ 1073ms | ✓ 885ms | http |
| 120.92.212.16:7890 | ✓ 952ms | ✓ 1178ms | 否 | ✓ 1250ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1651ms | ✓ 1194ms | ✓ 1137ms | 否 | 否 | http |
| 106.117.208.101:7890 | ✓ 999ms | ✓ 1466ms | ✓ 1133ms | ✓ 1923ms | ✓ 1220ms | http |
| 199.68.217.2:3128 | ✓ 821ms | ✓ 1313ms | ✓ 978ms | ✓ 892ms | ✓ 913ms | http |
| 185.114.73.2:1080 | ✓ 633ms | 否 | ✓ 1568ms | ✓ 1922ms | 否 | http |
| 103.94.251.91:8085 | ✓ 1436ms | 否 | ✓ 1613ms | ✓ 1582ms | 否 | http |
| 142.171.95.105:3128 | ✓ 775ms | ✓ 1840ms | ✓ 1434ms | ✓ 918ms | ✓ 612ms | http |
| 1.225.116.115:1080 | ✓ 1200ms | ✓ 1029ms | ✓ 793ms | ✓ 904ms | ✓ 702ms | http |
| 45.140.147.82:1082 | 否 | ✓ 1622ms | ✓ 1798ms | 否 | ✓ 1178ms | http |
| 38.145.218.9:8445 | ✓ 1425ms | ✓ 876ms | ✓ 1150ms | 否 | ✓ 800ms | http |
| 38.145.208.209:8447 | ✓ 1194ms | 否 | ✓ 844ms | ✓ 926ms | 否 | http |
| 38.34.179.213:8452 | ✓ 276ms | 否 | ✓ 197ms | ✓ 819ms | ✓ 730ms | http |
| 38.145.218.217:8450 | 否 | ✓ 825ms | ✓ 435ms | ✓ 854ms | ✓ 750ms | http |
| 45.136.131.37:8447 | 否 | ✓ 749ms | ✓ 307ms | ✓ 811ms | ✓ 730ms | http |
| 38.34.179.18:8451 | ✓ 326ms | ✓ 1619ms | ✓ 1413ms | ✓ 1103ms | ✓ 555ms | http |
| 45.136.131.26:8446 | ✓ 625ms | 否 | ✓ 150ms | ✓ 824ms | ✓ 934ms | http |
| 61.76.95.217:40088 | ✓ 1774ms | ✓ 1275ms | ✓ 1548ms | ✓ 1552ms | ✓ 1400ms | http |
| 114.231.72.199:1080 | ✓ 1671ms | 否 | ✓ 1272ms | ✓ 1148ms | ✓ 1321ms | http |
| 91.233.223.147:3128 | ✓ 1361ms | 否 | ✓ 1316ms | 否 | ✓ 1701ms | http |
| 185.255.178.231:3128 | ✓ 1337ms | 否 | ✓ 1694ms | 否 | ✓ 1677ms | http |
| 38.145.220.27:8445 | ✓ 679ms | ✓ 1889ms | ✓ 1848ms | 否 | 否 | http |
| 103.67.46.225:3125 | 否 | 否 | ✓ 1596ms | ✓ 1681ms | ✓ 1600ms | http |
| 38.145.218.162:8448 | ✓ 773ms | ✓ 1027ms | ✓ 447ms | ✓ 1005ms | ✓ 558ms | http |
| 38.145.218.160:8448 | ✓ 813ms | ✓ 1557ms | ✓ 305ms | ✓ 705ms | ✓ 626ms | http |
| 38.145.218.229:8449 | ✓ 872ms | ✓ 1182ms | ✓ 212ms | ✓ 701ms | ✓ 1114ms | http |
| 38.145.218.206:8444 | ✓ 1478ms | ✓ 943ms | ✓ 602ms | ✓ 990ms | ✓ 746ms | http |
| 38.34.183.13:8449 | ✓ 817ms | ✓ 1466ms | ✓ 487ms | ✓ 718ms | ✓ 690ms | http |
| 45.136.130.176:8451 | ✓ 784ms | ✓ 1156ms | ✓ 317ms | ✓ 990ms | ✓ 867ms | http |
| 38.145.208.227:8447 | ✓ 852ms | ✓ 1012ms | ✓ 1023ms | ✓ 767ms | ✓ 862ms | http |
| 38.145.220.173:8444 | ✓ 841ms | ✓ 1364ms | ✓ 90ms | ✓ 762ms | ✓ 1675ms | http |
| 38.34.179.21:8446 | ✓ 828ms | ✓ 1660ms | ✓ 1168ms | ✓ 936ms | ✓ 625ms | http |
| 45.136.131.34:8444 | ✓ 824ms | ✓ 1936ms | ✓ 735ms | ✓ 969ms | ✓ 1649ms | http |
| 45.136.131.29:8444 | ✓ 1441ms | ✓ 1167ms | ✓ 531ms | ✓ 1118ms | ✓ 1935ms | http |
| 38.145.220.11:8447 | ✓ 844ms | ✓ 1558ms | 否 | ✓ 1686ms | 否 | http |
| 121.230.8.247:1080 | ✓ 1272ms | ✓ 1359ms | ✓ 1126ms | 否 | 否 | http |
| 82.65.98.35:3128 | 否 | ✓ 1662ms | ✓ 707ms | ✓ 1749ms | ✓ 1287ms | http |
| 38.34.179.87:8447 | ✓ 1148ms | 否 | ✓ 986ms | ✓ 1361ms | 否 | http |
| 38.145.220.198:8446 | 否 | ✓ 1432ms | ✓ 690ms | ✓ 1352ms | ✓ 1678ms | http |
| 38.34.179.60:8450 | ✓ 347ms | 否 | ✓ 176ms | ✓ 693ms | ✓ 863ms | http |
| 38.34.183.47:8452 | 否 | ✓ 1693ms | ✓ 349ms | ✓ 699ms | ✓ 513ms | http |
| 38.34.179.69:8447 | 否 | ✓ 1422ms | ✓ 345ms | ✓ 821ms | ✓ 638ms | http |
| 38.34.179.186:8444 | 否 | 否 | ✓ 217ms | ✓ 801ms | ✓ 620ms | http |
| 38.145.208.229:8453 | 否 | ✓ 1952ms | ✓ 611ms | ✓ 1118ms | ✓ 930ms | http |
| 109.234.38.35:3128 | ✓ 739ms | 否 | 否 | ✓ 1492ms | ✓ 1169ms | http |
| 103.113.70.189:1081 | ✓ 403ms | 否 | ✓ 265ms | ✓ 1499ms | ✓ 1061ms | http |
| 139.198.113.42:10023 | ✓ 1081ms | ✓ 1699ms | ✓ 1499ms | ✓ 1967ms | ✓ 1067ms | http |
| 103.68.233.142:8097 | ✓ 1331ms | 否 | ✓ 1314ms | ✓ 1427ms | ✓ 1351ms | http |
| 47.251.25.177:443 | ✓ 1353ms | ✓ 1901ms | ✓ 284ms | ✓ 854ms | ✓ 694ms | http |

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
