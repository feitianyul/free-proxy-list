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

最后更新：2026-05-04 19:10:35 UTC（2026-05-05 03:10:35 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.71.229.255:3128 | ✓ 452ms | ✓ 1480ms | ✓ 1107ms | ✓ 1170ms | ✓ 829ms | http |
| 206.206.126.177:2412 | ✓ 885ms | ✓ 1545ms | ✓ 1254ms | ✓ 959ms | ✓ 761ms | http |
| 1.231.81.166:3128 | ✓ 1576ms | ✓ 1164ms | 否 | ✓ 1238ms | ✓ 1024ms | http |
| 113.160.132.26:8080 | ✓ 1639ms | ✓ 1892ms | ✓ 1054ms | 否 | ✓ 1325ms | http |
| 148.230.4.241:999 | ✓ 847ms | ✓ 1653ms | ✓ 871ms | ✓ 1827ms | 否 | http |
| 168.194.0.249:252 | ✓ 1412ms | 否 | ✓ 1569ms | 否 | ✓ 1822ms | http |
| 181.119.97.24:999 | ✓ 1568ms | 否 | ✓ 1770ms | ✓ 1957ms | 否 | http |
| 107.173.42.121:7890 | 否 | ✓ 1392ms | ✓ 729ms | ✓ 1731ms | ✓ 1570ms | http |
| 137.59.47.73:3128 | ✓ 1488ms | ✓ 1221ms | ✓ 1866ms | ✓ 1077ms | ✓ 1120ms | http |
| 80.92.204.47:1081 | ✓ 645ms | 否 | 否 | ✓ 1936ms | ✓ 1546ms | http |
| 120.92.212.16:7890 | ✓ 1116ms | 否 | ✓ 1350ms | ✓ 1201ms | ✓ 1875ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1103ms | ✓ 879ms | 否 | ✓ 1871ms | http |
| 47.85.51.197:1080 | ✓ 1410ms | ✓ 1235ms | ✓ 1028ms | ✓ 1642ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1172ms | 否 | ✓ 1789ms | ✓ 1236ms | ✓ 1113ms | http |
| 38.180.121.135:10808 | ✓ 756ms | ✓ 1963ms | 否 | 否 | ✓ 1300ms | http |
| 218.108.131.186:17890 | ✓ 627ms | ✓ 786ms | ✓ 604ms | ✓ 850ms | ✓ 682ms | http |
| 86.104.72.220:1081 | ✓ 1049ms | ✓ 1281ms | ✓ 895ms | 否 | ✓ 1343ms | http |
| 120.92.108.86:7890 | ✓ 1781ms | 否 | ✓ 1404ms | 否 | ✓ 1160ms | http |
| 45.153.231.229:8080 | ✓ 923ms | ✓ 1856ms | ✓ 1685ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 895ms | ✓ 916ms | 否 | 否 | ✓ 1404ms | http |
| 91.242.229.129:8092 | ✓ 862ms | ✓ 1582ms | ✓ 1085ms | ✓ 1715ms | 否 | http |
| 8.219.188.145:8118 | ✓ 1367ms | 否 | ✓ 1567ms | 否 | ✓ 1279ms | http |
| 185.191.236.162:3128 | ✓ 1310ms | ✓ 1801ms | ✓ 1970ms | 否 | ✓ 1828ms | http |
| 104.128.138.186:1080 | ✓ 1338ms | 否 | ✓ 1895ms | 否 | ✓ 1736ms | http |
| 34.101.184.164:3128 | ✓ 1706ms | 否 | 否 | ✓ 1582ms | ✓ 1412ms | http |
| 103.157.200.126:3128 | ✓ 1301ms | 否 | ✓ 1227ms | 否 | ✓ 1722ms | http |
| 43.133.44.89:8888 | ✓ 1894ms | 否 | ✓ 1704ms | ✓ 1247ms | 否 | http |
| 165.225.113.220:8800 | ✓ 700ms | 否 | 否 | ✓ 1199ms | ✓ 1063ms | http |
| 86.104.74.110:1082 | ✓ 1130ms | ✓ 1353ms | ✓ 1962ms | 否 | ✓ 1636ms | http |
| 86.104.74.110:1081 | ✓ 1133ms | ✓ 1406ms | ✓ 1904ms | 否 | 否 | http |
| 152.70.91.193:40000 | ✓ 1464ms | 否 | 否 | ✓ 1572ms | ✓ 1322ms | http |
| 154.64.232.35:8080 | ✓ 899ms | ✓ 1807ms | 否 | ✓ 755ms | 否 | http |
| 152.32.132.190:7890 | ✓ 645ms | ✓ 1739ms | 否 | ✓ 809ms | ✓ 1058ms | http |
| 89.208.106.138:10808 | ✓ 1278ms | ✓ 1673ms | ✓ 1317ms | 否 | 否 | http |
| 193.123.250.39:1080 | 否 | 否 | ✓ 1761ms | ✓ 1464ms | ✓ 1996ms | http |
| 62.113.119.14:8080 | ✓ 783ms | ✓ 1703ms | ✓ 759ms | ✓ 1679ms | ✓ 1224ms | http |
| 171.80.137.18:7890 | ✓ 703ms | ✓ 988ms | ✓ 724ms | ✓ 1132ms | ✓ 809ms | http |
| 116.171.106.111:3443 | ✓ 1241ms | 否 | ✓ 1304ms | ✓ 1383ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1845ms | ✓ 1308ms | ✓ 1144ms | ✓ 1128ms | ✓ 1232ms | http |
| 190.12.150.244:999 | ✓ 1696ms | ✓ 1814ms | ✓ 1052ms | 否 | 否 | http |
| 86.104.72.219:1081 | ✓ 666ms | ✓ 1514ms | ✓ 841ms | ✓ 1735ms | 否 | http |
| 116.171.106.26:3443 | ✓ 1425ms | 否 | ✓ 1208ms | ✓ 1446ms | ✓ 1233ms | http |
| 47.77.216.82:1080 | ✓ 34ms | ✓ 750ms | ✓ 757ms | ✓ 958ms | ✓ 629ms | http |
| 8.219.97.248:80 | 否 | ✓ 1813ms | ✓ 1907ms | ✓ 1974ms | 否 | http |
| 220.197.44.36:3128 | ✓ 1517ms | 否 | ✓ 1725ms | ✓ 1582ms | 否 | http |
| 45.125.67.37:8443 | ✓ 817ms | 否 | ✓ 809ms | ✓ 979ms | ✓ 1008ms | http |
| 106.10.55.212:1121 | ✓ 522ms | ✓ 1122ms | 否 | ✓ 1288ms | ✓ 1700ms | http |
| 118.113.246.73:1080 | ✓ 1075ms | ✓ 1404ms | ✓ 1133ms | ✓ 1547ms | ✓ 1124ms | http |
| 47.105.98.23:3128 | ✓ 921ms | ✓ 1582ms | ✓ 1035ms | ✓ 1402ms | ✓ 1020ms | http |
| 103.176.97.196:8082 | ✓ 1968ms | 否 | ✓ 1201ms | ✓ 1440ms | ✓ 1370ms | http |
| 129.213.162.27:17777 | 否 | ✓ 1945ms | 否 | ✓ 1663ms | ✓ 1249ms | http |
| 101.6.65.112:10080 | ✓ 1297ms | ✓ 1472ms | ✓ 1385ms | ✓ 1752ms | ✓ 1188ms | http |
| 3.101.133.120:80 | ✓ 195ms | ✓ 1027ms | ✓ 1262ms | ✓ 886ms | ✓ 816ms | http |
| 43.99.54.236:5555 | ✓ 648ms | ✓ 871ms | ✓ 677ms | ✓ 788ms | ✓ 627ms | http |
| 103.35.190.69:1082 | ✓ 763ms | ✓ 1478ms | ✓ 535ms | ✓ 1404ms | ✓ 922ms | http |
| 45.59.122.132:80 | ✓ 1308ms | 否 | ✓ 1327ms | 否 | ✓ 1736ms | http |
| 86.104.72.220:1082 | ✓ 463ms | ✓ 1173ms | ✓ 595ms | ✓ 1301ms | 否 | http |
| 104.248.195.47:8080 | ✓ 1160ms | ✓ 1721ms | ✓ 1820ms | ✓ 1772ms | ✓ 1553ms | http |
| 38.188.247.12:999 | ✓ 880ms | ✓ 1987ms | ✓ 1477ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 1002ms | 否 | ✓ 1378ms | 否 | ✓ 1742ms | http |
| 61.52.131.172:8443 | ✓ 737ms | ✓ 995ms | ✓ 806ms | ✓ 922ms | ✓ 766ms | http |
| 91.238.105.64:2024 | ✓ 981ms | 否 | ✓ 1906ms | 否 | ✓ 1937ms | http |
| 194.59.247.34:10808 | ✓ 1306ms | ✓ 1993ms | ✓ 1810ms | 否 | ✓ 1773ms | http |
| 103.172.70.173:8080 | ✓ 1225ms | ✓ 1922ms | ✓ 1615ms | ✓ 1320ms | ✓ 1334ms | http |
| 103.156.14.165:8080 | ✓ 1933ms | 否 | ✓ 1222ms | 否 | ✓ 1326ms | http |
| 103.39.51.207:8080 | ✓ 1370ms | 否 | 否 | ✓ 1965ms | ✓ 1447ms | http |
| 115.191.40.196:7890 | 否 | ✓ 996ms | ✓ 809ms | 否 | ✓ 1702ms | http |

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
