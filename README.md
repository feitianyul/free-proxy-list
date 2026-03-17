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

最后更新：2026-03-17 17:58:13 UTC（2026-03-18 01:58:13 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 101.47.73.135:3128 | ✓ 1104ms | 否 | ✓ 1114ms | ✓ 1521ms | ✓ 1316ms | http |
| 202.155.12.161:443 | ✓ 1469ms | 否 | ✓ 1053ms | ✓ 1195ms | ✓ 1110ms | http |
| 147.161.239.240:8800 | ✓ 1027ms | ✓ 1791ms | ✓ 1283ms | ✓ 1668ms | ✓ 1428ms | http |
| 147.161.210.140:8800 | ✓ 1468ms | 否 | ✓ 1186ms | ✓ 1092ms | ✓ 1461ms | http |
| 113.160.132.26:8080 | ✓ 1671ms | 否 | ✓ 1136ms | ✓ 1831ms | ✓ 1444ms | http |
| 45.167.124.52:8080 | ✓ 1431ms | 否 | ✓ 929ms | ✓ 1979ms | ✓ 1482ms | http |
| 183.249.5.117:22222 | ✓ 1213ms | ✓ 1005ms | ✓ 876ms | ✓ 1321ms | ✓ 841ms | http |
| 1.231.81.166:3128 | ✓ 815ms | ✓ 1037ms | ✓ 981ms | ✓ 1193ms | ✓ 980ms | http |
| 137.220.150.170:6005 | 否 | 否 | ✓ 1047ms | ✓ 1312ms | ✓ 1055ms | http |
| 35.225.22.61:80 | ✓ 594ms | ✓ 1607ms | ✓ 1183ms | 否 | ✓ 974ms | http |
| 222.184.48.251:22222 | 否 | 否 | ✓ 1075ms | ✓ 1444ms | ✓ 1325ms | http |
| 104.129.203.245:10733 | 否 | ✓ 1161ms | ✓ 1230ms | ✓ 1037ms | ✓ 720ms | http |
| 104.129.203.245:10026 | 否 | ✓ 1136ms | ✓ 1230ms | ✓ 1070ms | ✓ 708ms | http |
| 104.129.203.245:10139 | 否 | ✓ 1815ms | ✓ 218ms | ✓ 960ms | ✓ 722ms | http |
| 138.124.53.25:7443 | ✓ 924ms | 否 | ✓ 1203ms | 否 | ✓ 1179ms | http |
| 101.43.127.100:8877 | ✓ 1236ms | ✓ 1489ms | ✓ 1062ms | ✓ 1327ms | ✓ 1795ms | http |
| 137.220.150.152:6005 | ✓ 952ms | 否 | ✓ 1083ms | ✓ 1293ms | ✓ 964ms | http |
| 183.249.5.105:22222 | ✓ 920ms | ✓ 1019ms | ✓ 889ms | ✓ 1093ms | ✓ 1000ms | http |
| 183.249.5.111:22222 | ✓ 805ms | ✓ 1284ms | ✓ 867ms | ✓ 1113ms | ✓ 845ms | http |
| 193.23.200.251:10808 | ✓ 1259ms | 否 | ✓ 1353ms | ✓ 1693ms | ✓ 1302ms | http |
| 45.136.198.40:3128 | ✓ 1241ms | ✓ 1798ms | ✓ 1540ms | 否 | ✓ 1921ms | http |
| 103.113.70.189:1081 | ✓ 1361ms | ✓ 980ms | 否 | 否 | ✓ 1768ms | http |
| 38.34.179.60:8450 | ✓ 952ms | ✓ 1224ms | ✓ 1222ms | ✓ 1973ms | ✓ 1548ms | http |
| 120.92.212.16:7890 | ✓ 1088ms | 否 | ✓ 1390ms | ✓ 1398ms | ✓ 1091ms | http |
| 137.220.150.104:6005 | ✓ 1725ms | 否 | ✓ 1568ms | ✓ 1528ms | ✓ 1308ms | http |
| 86.53.183.16:1080 | ✓ 1416ms | 否 | ✓ 1124ms | ✓ 1737ms | ✓ 1397ms | http |
| 77.110.113.24:40000 | ✓ 1503ms | 否 | ✓ 880ms | 否 | ✓ 1835ms | http |
| 168.235.110.63:3128 | ✓ 402ms | ✓ 958ms | ✓ 1018ms | ✓ 1128ms | ✓ 947ms | http |
| 38.145.218.160:8448 | ✓ 1553ms | 否 | ✓ 1464ms | 否 | ✓ 1215ms | http |
| 59.46.216.131:30001 | ✓ 1030ms | 否 | ✓ 1237ms | 否 | ✓ 1235ms | http |
| 137.220.151.110:6005 | ✓ 882ms | 否 | ✓ 1474ms | ✓ 1342ms | ✓ 1242ms | http |
| 47.77.193.180:1080 | ✓ 1477ms | ✓ 978ms | ✓ 294ms | ✓ 829ms | ✓ 617ms | http |
| 219.117.204.211:7799 | ✓ 1743ms | 否 | ✓ 1145ms | ✓ 959ms | ✓ 995ms | http |
| 120.92.212.16:8890 | ✓ 1213ms | ✓ 1387ms | ✓ 1338ms | 否 | ✓ 1079ms | http |
| 88.80.150.82:8080 | ✓ 1067ms | 否 | 否 | ✓ 1929ms | ✓ 1828ms | https |
| 45.136.130.173:8448 | ✓ 958ms | ✓ 974ms | ✓ 497ms | ✓ 1283ms | ✓ 808ms | http |
| 45.136.130.168:8448 | ✓ 957ms | ✓ 975ms | ✓ 496ms | ✓ 1294ms | ✓ 790ms | http |
| 192.71.213.85:9091 | ✓ 1083ms | 否 | ✓ 1506ms | ✓ 1727ms | 否 | http |
| 91.107.148.58:53967 | ✓ 968ms | 否 | 否 | ✓ 1588ms | ✓ 1793ms | http |
| 202.141.161.53:10808 | 否 | ✓ 1413ms | ✓ 1356ms | ✓ 1459ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1633ms | 否 | 否 | ✓ 1274ms | ✓ 972ms | http |
| 103.82.23.118:5171 | ✓ 1493ms | 否 | ✓ 1468ms | 否 | ✓ 1624ms | http |
| 222.109.119.178:3128 | 否 | 否 | ✓ 987ms | ✓ 1306ms | ✓ 1815ms | http |
| 45.136.131.36:8450 | ✓ 749ms | ✓ 1469ms | 否 | ✓ 1101ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1691ms | ✓ 1814ms | ✓ 1781ms | 否 | 否 | http |
| 194.5.212.40:8080 | ✓ 584ms | ✓ 1862ms | ✓ 871ms | 否 | ✓ 1498ms | http |
| 217.77.102.18:3128 | ✓ 1489ms | ✓ 1961ms | ✓ 1164ms | ✓ 1836ms | ✓ 1591ms | http |
| 103.39.51.190:8080 | ✓ 1865ms | 否 | 否 | ✓ 1649ms | ✓ 1472ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 1754ms | ✓ 1139ms | ✓ 1654ms | http |
| 103.117.100.127:13082 | ✓ 935ms | ✓ 1773ms | ✓ 740ms | ✓ 980ms | ✓ 764ms | http |
| 210.76.192.9:10810 | ✓ 1053ms | ✓ 1871ms | ✓ 1350ms | ✓ 1871ms | ✓ 1076ms | http |
| 121.230.8.208:1080 | ✓ 1578ms | ✓ 1752ms | ✓ 1403ms | ✓ 1737ms | ✓ 1267ms | http |
| 38.34.179.89:8448 | ✓ 1084ms | ✓ 944ms | ✓ 316ms | ✓ 931ms | ✓ 754ms | http |
| 38.34.179.103:8450 | ✓ 624ms | ✓ 1025ms | ✓ 698ms | ✓ 999ms | ✓ 767ms | http |
| 38.34.179.18:8448 | ✓ 1398ms | ✓ 941ms | ✓ 559ms | 否 | 否 | http |
| 38.34.179.213:8450 | ✓ 370ms | ✓ 969ms | ✓ 895ms | ✓ 1300ms | ✓ 764ms | http |
| 38.34.179.7:8448 | 否 | 否 | ✓ 354ms | ✓ 1173ms | ✓ 1341ms | http |
| 38.145.203.135:8453 | ✓ 415ms | ✓ 967ms | ✓ 552ms | ✓ 1242ms | ✓ 740ms | http |
| 38.34.178.155:8448 | ✓ 1481ms | ✓ 879ms | ✓ 405ms | ✓ 961ms | ✓ 767ms | http |
| 62.60.177.204:34094 | ✓ 380ms | ✓ 1281ms | ✓ 1118ms | ✓ 1999ms | ✓ 823ms | http |
| 183.249.5.110:22222 | ✓ 855ms | ✓ 1085ms | ✓ 861ms | ✓ 1068ms | ✓ 1016ms | http |
| 14.143.222.113:10155 | ✓ 1720ms | 否 | ✓ 1120ms | ✓ 1545ms | 否 | http |

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
