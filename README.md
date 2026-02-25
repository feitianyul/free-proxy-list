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

最后更新：2026-02-25 15:44:01 UTC（2026-02-25 23:44:01 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 35.225.22.61:80 | 否 | 否 | ✓ 1633ms | ✓ 1003ms | ✓ 916ms | http |
| 51.81.46.174:3128 | ✓ 314ms | 否 | ✓ 1195ms | ✓ 1707ms | ✓ 1138ms | http |
| 120.92.212.16:7890 | ✓ 1012ms | ✓ 1222ms | ✓ 1008ms | ✓ 1254ms | ✓ 1001ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1483ms | ✓ 1193ms | ✓ 1515ms | ✓ 1208ms | http |
| 20.210.39.153:8561 | ✓ 617ms | 否 | ✓ 799ms | ✓ 890ms | ✓ 634ms | http |
| 20.78.118.91:8561 | ✓ 625ms | 否 | ✓ 794ms | ✓ 888ms | ✓ 632ms | http |
| 20.78.26.206:8561 | ✓ 617ms | 否 | ✓ 798ms | ✓ 890ms | ✓ 631ms | http |
| 202.152.44.18:8081 | ✓ 1437ms | 否 | ✓ 1102ms | ✓ 1377ms | ✓ 1174ms | http |
| 165.227.5.10:8888 | ✓ 607ms | 否 | ✓ 623ms | 否 | ✓ 1218ms | http |
| 178.130.47.129:1082 | ✓ 718ms | 否 | ✓ 1016ms | ✓ 986ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1678ms | ✓ 1865ms | 否 | ✓ 1396ms | ✓ 802ms | http |
| 202.152.44.19:8081 | 否 | 否 | ✓ 919ms | ✓ 1820ms | ✓ 1010ms | http |
| 116.80.64.158:7777 | ✓ 1726ms | 否 | ✓ 1559ms | 否 | ✓ 1677ms | http |
| 85.208.108.43:2094 | 否 | 否 | ✓ 1424ms | ✓ 1102ms | ✓ 973ms | http |
| 188.166.208.168:9876 | ✓ 782ms | 否 | ✓ 1069ms | ✓ 1204ms | ✓ 940ms | http |
| 101.43.255.96:80 | ✓ 1771ms | ✓ 1922ms | ✓ 1230ms | 否 | ✓ 1017ms | http |
| 81.70.169.194:80 | ✓ 960ms | ✓ 1261ms | ✓ 1296ms | ✓ 1255ms | ✓ 1850ms | http |
| 45.140.147.155:1081 | ✓ 651ms | ✓ 1738ms | ✓ 1194ms | 否 | ✓ 1409ms | http |
| 217.217.254.94:8080 | ✓ 936ms | 否 | ✓ 1290ms | 否 | ✓ 1413ms | http |
| 121.230.8.158:1080 | ✓ 1089ms | 否 | ✓ 1331ms | ✓ 1802ms | 否 | http |
| 103.84.95.54:7890 | ✓ 674ms | 否 | 否 | ✓ 1177ms | ✓ 1146ms | http |
| 120.46.152.136:3128 | ✓ 1333ms | 否 | ✓ 1296ms | ✓ 1416ms | ✓ 941ms | http |
| 35.234.17.221:8080 | ✓ 952ms | 否 | 否 | ✓ 1235ms | ✓ 939ms | http |
| 45.136.198.40:3128 | ✓ 997ms | 否 | ✓ 1630ms | 否 | ✓ 1791ms | http |
| 36.147.78.166:80 | 否 | ✓ 1728ms | ✓ 1807ms | 否 | ✓ 1447ms | http |
| 117.86.6.194:1080 | ✓ 1706ms | 否 | ✓ 1048ms | 否 | ✓ 910ms | http |
| 109.107.181.253:1080 | ✓ 1251ms | 否 | 否 | ✓ 1839ms | ✓ 1848ms | http |
| 124.16.93.70:7890 | ✓ 899ms | ✓ 1438ms | ✓ 1043ms | ✓ 1209ms | ✓ 959ms | http |
| 45.88.0.98:3128 | ✓ 640ms | ✓ 1435ms | ✓ 1764ms | ✓ 1876ms | 否 | http |
| 62.113.119.14:8080 | ✓ 805ms | 否 | ✓ 789ms | ✓ 1858ms | ✓ 1414ms | http |
| 45.88.0.115:3128 | ✓ 568ms | ✓ 1594ms | 否 | 否 | ✓ 1225ms | http |
| 45.88.0.113:3128 | ✓ 681ms | 否 | ✓ 1181ms | ✓ 1817ms | 否 | http |
| 45.88.0.111:3128 | ✓ 1099ms | 否 | ✓ 1757ms | ✓ 1454ms | ✓ 1175ms | http |
| 45.88.0.114:3128 | ✓ 1105ms | 否 | ✓ 1751ms | ✓ 1463ms | ✓ 1177ms | http |
| 45.88.0.99:3128 | ✓ 1103ms | 否 | ✓ 1753ms | ✓ 1458ms | ✓ 1180ms | http |
| 45.88.0.116:3128 | ✓ 1103ms | ✓ 1819ms | ✓ 1933ms | ✓ 1440ms | ✓ 1160ms | http |
| 45.88.0.117:3128 | ✓ 1099ms | 否 | ✓ 1756ms | ✓ 1462ms | ✓ 1178ms | http |
| 113.45.250.180:443 | ✓ 974ms | ✓ 1256ms | ✓ 1042ms | ✓ 1185ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1009ms | 否 | ✓ 1223ms | 否 | ✓ 1025ms | http |
| 81.177.48.54:2080 | ✓ 1508ms | 否 | ✓ 1320ms | ✓ 1956ms | ✓ 1475ms | http |
| 103.69.84.106:3131 | 否 | 否 | ✓ 1401ms | ✓ 1415ms | ✓ 938ms | http |
| 103.39.51.190:8080 | ✓ 1466ms | 否 | 否 | ✓ 1566ms | ✓ 1341ms | http |
| 187.216.141.46:3128 | 否 | ✓ 1216ms | ✓ 1282ms | ✓ 1574ms | 否 | http |
| 172.212.68.37:3128 | ✓ 866ms | 否 | ✓ 1462ms | ✓ 1451ms | ✓ 1213ms | http |
| 36.147.78.166:443 | ✓ 1703ms | ✓ 1746ms | ✓ 1705ms | ✓ 1829ms | ✓ 1636ms | http |
| 121.204.158.249:3128 | 否 | 否 | ✓ 1371ms | ✓ 1331ms | ✓ 1025ms | http |
| 137.184.6.117:3128 | ✓ 1353ms | ✓ 920ms | ✓ 362ms | ✓ 997ms | ✓ 643ms | http |
| 71.204.156.52:443 | ✓ 1837ms | ✓ 1965ms | ✓ 874ms | ✓ 1058ms | ✓ 1222ms | http |
| 103.149.99.128:3128 | ✓ 1669ms | ✓ 1625ms | ✓ 1618ms | ✓ 1283ms | 否 | http |

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
