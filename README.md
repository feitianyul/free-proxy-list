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

最后更新：2026-03-09 15:52:46 UTC（2026-03-09 23:52:46 UTC+8）

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
| 61.72.110.114:3128 | 否 | 否 | ✓ 1624ms | ✓ 1245ms | ✓ 1930ms | http |
| 211.171.114.154:3128 | ✓ 1170ms | 否 | ✓ 1894ms | ✓ 1610ms | ✓ 1173ms | http |
| 1.231.81.166:3128 | ✓ 1862ms | ✓ 1619ms | ✓ 1626ms | ✓ 1382ms | ✓ 1181ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1095ms | ✓ 1861ms | ✓ 1342ms | http |
| 35.225.22.61:80 | ✓ 531ms | 否 | ✓ 135ms | ✓ 980ms | ✓ 959ms | http |
| 136.49.34.18:8888 | ✓ 691ms | 否 | ✓ 1792ms | ✓ 1981ms | 否 | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1304ms | ✓ 1454ms | ✓ 1150ms | http |
| 120.92.212.16:8890 | ✓ 1333ms | 否 | ✓ 1041ms | ✓ 1428ms | ✓ 1097ms | http |
| 168.235.110.63:3128 | ✓ 646ms | ✓ 1828ms | ✓ 961ms | 否 | ✓ 843ms | http |
| 34.96.238.40:8080 | 否 | 否 | ✓ 1329ms | ✓ 1328ms | ✓ 967ms | http |
| 39.104.201.40:7890 | ✓ 1023ms | 否 | ✓ 1776ms | 否 | ✓ 1066ms | http |
| 178.236.245.59:3128 | ✓ 1951ms | 否 | ✓ 1687ms | 否 | ✓ 1796ms | http |
| 81.70.169.194:80 | 否 | ✓ 1454ms | ✓ 1419ms | 否 | ✓ 1113ms | http |
| 101.43.255.96:80 | ✓ 1172ms | 否 | ✓ 1038ms | ✓ 1653ms | 否 | http |
| 190.9.109.207:999 | ✓ 910ms | ✓ 1431ms | ✓ 1162ms | ✓ 1213ms | ✓ 1070ms | http |
| 190.9.109.198:999 | ✓ 912ms | ✓ 1449ms | ✓ 1018ms | ✓ 1354ms | 否 | http |
| 202.155.12.161:443 | ✓ 1403ms | 否 | ✓ 1477ms | ✓ 1027ms | ✓ 909ms | http |
| 45.129.141.143:3128 | ✓ 1205ms | 否 | ✓ 1693ms | ✓ 1909ms | ✓ 1447ms | http |
| 150.230.211.74:8080 | ✓ 1999ms | 否 | ✓ 1874ms | ✓ 1400ms | ✓ 1003ms | http |
| 165.227.5.10:8888 | 否 | ✓ 1539ms | 否 | ✓ 1093ms | ✓ 1242ms | http |
| 45.136.198.40:3128 | ✓ 694ms | 否 | ✓ 1478ms | ✓ 1519ms | ✓ 1144ms | http |
| 8.219.97.248:80 | ✓ 1650ms | 否 | ✓ 1117ms | ✓ 1964ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1225ms | ✓ 1555ms | ✓ 1566ms | ✓ 1643ms | ✓ 1289ms | http |
| 194.213.18.200:443 | ✓ 1190ms | ✓ 1845ms | 否 | 否 | ✓ 1796ms | http |
| 138.124.53.25:7443 | ✓ 1390ms | 否 | 否 | ✓ 1974ms | ✓ 1520ms | http |
| 45.140.147.82:1082 | ✓ 530ms | ✓ 1476ms | ✓ 1340ms | 否 | ✓ 1252ms | http |
| 45.140.147.82:1081 | ✓ 1629ms | ✓ 1277ms | ✓ 481ms | 否 | ✓ 1246ms | http |
| 61.72.110.54:3128 | ✓ 770ms | 否 | ✓ 926ms | ✓ 1174ms | ✓ 954ms | http |
| 61.72.221.94:3128 | ✓ 754ms | ✓ 1277ms | ✓ 889ms | ✓ 1142ms | ✓ 927ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1838ms | ✓ 989ms | ✓ 1209ms | ✓ 912ms | http |
| 125.128.12.14:3128 | 否 | ✓ 1981ms | ✓ 981ms | ✓ 1221ms | ✓ 906ms | http |
| 5.101.0.233:3128 | ✓ 1233ms | ✓ 1656ms | ✓ 1311ms | ✓ 1874ms | ✓ 1492ms | http |
| 194.147.90.23:3128 | ✓ 1228ms | 否 | ✓ 1488ms | ✓ 1921ms | ✓ 1753ms | http |
| 14.225.212.37:7890 | 否 | ✓ 1665ms | ✓ 935ms | 否 | ✓ 1158ms | http |
| 120.92.212.16:7890 | ✓ 1049ms | ✓ 1384ms | 否 | 否 | ✓ 1084ms | http |
| 116.80.49.161:3172 | 否 | 否 | ✓ 1674ms | ✓ 1982ms | ✓ 1840ms | http |
| 121.237.181.137:8888 | ✓ 1022ms | ✓ 1675ms | ✓ 1704ms | 否 | ✓ 1364ms | http |
| 14.56.107.244:3128 | ✓ 1603ms | ✓ 1762ms | ✓ 1583ms | ✓ 1401ms | ✓ 911ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1366ms | ✓ 1366ms | ✓ 1400ms | 否 | http |
| 192.71.213.85:9091 | ✓ 859ms | 否 | ✓ 854ms | ✓ 1886ms | 否 | http |
| 205.209.118.30:3138 | ✓ 674ms | ✓ 1791ms | 否 | 否 | ✓ 1925ms | http |
| 210.77.29.245:7890 | ✓ 1062ms | ✓ 1589ms | ✓ 1139ms | ✓ 1323ms | ✓ 1080ms | http |
| 121.230.9.4:1080 | ✓ 1465ms | ✓ 1825ms | ✓ 1290ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1656ms | 否 | ✓ 1597ms | ✓ 1587ms | ✓ 1270ms | http |
| 118.113.246.72:1080 | 否 | ✓ 1762ms | ✓ 1894ms | 否 | ✓ 1530ms | http |
| 103.39.51.190:8080 | ✓ 1788ms | 否 | ✓ 1989ms | 否 | ✓ 1585ms | http |
| 152.42.213.210:80 | ✓ 1633ms | 否 | 否 | ✓ 1226ms | ✓ 971ms | http |
| 172.212.68.37:3128 | ✓ 1082ms | 否 | ✓ 1049ms | 否 | ✓ 983ms | http |
| 47.77.193.180:1080 | ✓ 335ms | ✓ 1101ms | ✓ 1925ms | 否 | 否 | http |

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
