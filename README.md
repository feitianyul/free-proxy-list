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

最后更新：2026-05-04 15:05:01 UTC（2026-05-04 23:05:01 UTC+8）

**代理总数：66**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 826ms | ✓ 1970ms | ✓ 1018ms | ✓ 1192ms | ✓ 1217ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1577ms | 否 | ✓ 1449ms | ✓ 1202ms | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1670ms | ✓ 1775ms | ✓ 1456ms | http |
| 181.119.97.24:999 | ✓ 1894ms | 否 | ✓ 1840ms | 否 | ✓ 1780ms | http |
| 107.173.42.121:7890 | 否 | ✓ 1180ms | ✓ 578ms | ✓ 1299ms | 否 | http |
| 206.206.126.177:2412 | ✓ 907ms | 否 | ✓ 963ms | ✓ 1242ms | ✓ 992ms | http |
| 194.59.247.34:10808 | ✓ 460ms | ✓ 1557ms | 否 | 否 | ✓ 1365ms | http |
| 218.108.131.186:17890 | 否 | ✓ 1440ms | ✓ 1059ms | ✓ 1310ms | ✓ 1043ms | http |
| 20.27.11.248:8561 | 否 | ✓ 1115ms | ✓ 931ms | ✓ 1300ms | ✓ 1234ms | http |
| 20.27.14.220:8561 | 否 | ✓ 1419ms | ✓ 828ms | ✓ 1265ms | ✓ 1211ms | http |
| 20.27.15.111:8561 | 否 | ✓ 1775ms | ✓ 880ms | ✓ 1298ms | ✓ 1229ms | http |
| 20.27.13.35:8561 | 否 | 否 | ✓ 956ms | ✓ 1354ms | ✓ 1335ms | http |
| 217.76.245.80:999 | ✓ 786ms | ✓ 1444ms | ✓ 1347ms | ✓ 1437ms | ✓ 1423ms | http |
| 47.85.51.197:1080 | 否 | ✓ 1988ms | ✓ 688ms | 否 | ✓ 664ms | http |
| 190.12.150.244:999 | ✓ 873ms | ✓ 1859ms | ✓ 1526ms | ✓ 1960ms | ✓ 1831ms | http |
| 20.27.15.49:8561 | ✓ 1651ms | ✓ 1460ms | ✓ 871ms | ✓ 1248ms | ✓ 980ms | http |
| 20.210.76.175:8561 | ✓ 1651ms | ✓ 1347ms | ✓ 980ms | ✓ 1263ms | ✓ 971ms | http |
| 20.210.76.178:8561 | ✓ 1653ms | ✓ 1482ms | ✓ 850ms | ✓ 1255ms | ✓ 984ms | http |
| 20.210.76.104:8561 | ✓ 1653ms | ✓ 1750ms | ✓ 901ms | ✓ 1124ms | ✓ 1098ms | http |
| 46.105.190.40:3128 | ✓ 501ms | 否 | ✓ 697ms | ✓ 1673ms | ✓ 1283ms | http |
| 103.35.190.69:1082 | ✓ 123ms | 否 | ✓ 152ms | ✓ 1305ms | ✓ 817ms | http |
| 120.92.212.16:8890 | ✓ 1253ms | 否 | ✓ 1967ms | ✓ 1998ms | 否 | http |
| 47.77.216.82:1080 | ✓ 1362ms | 否 | ✓ 286ms | ✓ 900ms | 否 | http |
| 154.64.232.35:8080 | ✓ 751ms | 否 | ✓ 1143ms | 否 | ✓ 1395ms | http |
| 80.92.204.47:1082 | ✓ 1581ms | 否 | ✓ 1483ms | 否 | ✓ 1679ms | http |
| 62.113.119.14:8080 | ✓ 682ms | 否 | ✓ 648ms | ✓ 1853ms | ✓ 1554ms | http |
| 45.153.231.229:8080 | ✓ 780ms | ✓ 1880ms | ✓ 1300ms | 否 | ✓ 1946ms | http |
| 8.219.97.248:80 | ✓ 1522ms | 否 | 否 | ✓ 1903ms | ✓ 1535ms | http |
| 130.61.174.200:1080 | 否 | ✓ 1562ms | ✓ 1502ms | ✓ 1278ms | 否 | http |
| 103.157.200.126:3128 | ✓ 1611ms | 否 | ✓ 1043ms | ✓ 1680ms | ✓ 1291ms | http |
| 94.131.118.129:1081 | ✓ 979ms | ✓ 1515ms | ✓ 605ms | ✓ 1725ms | ✓ 1230ms | http |
| 103.35.190.69:1081 | ✓ 1246ms | ✓ 986ms | ✓ 63ms | ✓ 1988ms | ✓ 1907ms | http |
| 20.78.26.206:8561 | ✓ 1795ms | 否 | ✓ 879ms | ✓ 1025ms | ✓ 901ms | http |
| 20.78.118.91:8561 | ✓ 1789ms | 否 | ✓ 882ms | ✓ 1020ms | ✓ 920ms | http |
| 20.210.39.153:8561 | ✓ 1787ms | ✓ 1157ms | ✓ 709ms | ✓ 1127ms | ✓ 1007ms | http |
| 20.18.193.135:8561 | ✓ 1787ms | 否 | ✓ 1499ms | ✓ 1999ms | ✓ 1895ms | http |
| 45.59.122.132:80 | ✓ 716ms | 否 | ✓ 860ms | ✓ 1609ms | ✓ 1460ms | http |
| 45.125.67.37:8443 | ✓ 1145ms | 否 | ✓ 1260ms | ✓ 1259ms | 否 | http |
| 152.32.132.190:7890 | ✓ 881ms | 否 | 否 | ✓ 1414ms | ✓ 1792ms | http |
| 101.32.244.83:8080 | ✓ 1232ms | ✓ 1990ms | ✓ 1151ms | ✓ 1679ms | ✓ 1584ms | http |
| 121.43.196.213:8222 | ✓ 1136ms | ✓ 1258ms | ✓ 1092ms | ✓ 1395ms | ✓ 1084ms | http |
| 121.43.196.210:8222 | ✓ 1140ms | ✓ 1284ms | ✓ 1061ms | ✓ 1360ms | ✓ 1111ms | http |
| 34.101.184.164:3128 | ✓ 1875ms | 否 | ✓ 1702ms | 否 | ✓ 1172ms | http |
| 211.95.152.50:45046 | ✓ 1789ms | 否 | ✓ 1638ms | 否 | ✓ 1327ms | http |
| 177.93.33.55:999 | ✓ 1124ms | 否 | 否 | ✓ 1823ms | ✓ 1680ms | http |
| 86.104.72.219:1081 | ✓ 504ms | ✓ 1625ms | ✓ 67ms | ✓ 1364ms | ✓ 753ms | http |
| 86.104.72.219:1082 | ✓ 515ms | ✓ 932ms | ✓ 189ms | ✓ 1027ms | ✓ 890ms | http |
| 89.208.106.138:10808 | ✓ 1646ms | ✓ 1346ms | ✓ 1354ms | 否 | 否 | http |
| 193.123.250.39:1080 | ✓ 1639ms | 否 | ✓ 1450ms | ✓ 1130ms | ✓ 1067ms | http |
| 3.101.133.120:80 | ✓ 379ms | 否 | 否 | ✓ 1534ms | ✓ 951ms | http |
| 81.26.190.143:1080 | ✓ 1014ms | 否 | ✓ 1915ms | 否 | ✓ 1686ms | http |
| 103.82.23.118:5182 | ✓ 1629ms | 否 | ✓ 1582ms | 否 | ✓ 1798ms | http |
| 148.230.4.241:999 | ✓ 850ms | ✓ 1653ms | ✓ 812ms | 否 | 否 | http |
| 80.92.204.47:1081 | ✓ 474ms | 否 | ✓ 892ms | ✓ 1943ms | ✓ 1229ms | http |
| 77.232.142.164:3128 | ✓ 553ms | ✓ 1675ms | ✓ 1156ms | 否 | ✓ 1742ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1412ms | ✓ 1918ms | ✓ 1913ms | http |
| 38.180.62.47:10808 | ✓ 1273ms | 否 | ✓ 1311ms | ✓ 1659ms | ✓ 1312ms | http |
| 86.104.72.220:1081 | ✓ 1004ms | ✓ 1271ms | ✓ 85ms | ✓ 1012ms | ✓ 810ms | http |
| 45.173.12.140:1994 | ✓ 1835ms | 否 | ✓ 1176ms | ✓ 1559ms | ✓ 1271ms | http |
| 200.125.171.254:999 | 否 | 否 | ✓ 1133ms | ✓ 1770ms | ✓ 1070ms | http |
| 121.230.9.11:1080 | 否 | ✓ 1512ms | ✓ 1269ms | 否 | ✓ 1280ms | http |
| 103.18.77.24:1111 | ✓ 1771ms | 否 | ✓ 1834ms | ✓ 1959ms | 否 | http |
| 157.254.37.238:999 | ✓ 1272ms | 否 | ✓ 1499ms | 否 | ✓ 1272ms | http |
| 61.52.131.172:8443 | ✓ 1127ms | ✓ 1445ms | ✓ 1253ms | ✓ 1526ms | ✓ 1198ms | http |
| 129.213.162.27:17777 | ✓ 958ms | 否 | 否 | ✓ 1632ms | ✓ 1213ms | http |
| 1.231.81.166:3128 | ✓ 1249ms | 否 | ✓ 1948ms | ✓ 1414ms | ✓ 1018ms | http |

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
